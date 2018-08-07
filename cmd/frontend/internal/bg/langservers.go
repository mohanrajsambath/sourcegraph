package bg

import (
	"context"

	log15 "gopkg.in/inconshreveable/log15.v2"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/internal/pkg/langservers"
	"github.com/sourcegraph/sourcegraph/pkg/conf"
)

// StartLangServers should be invoked on startup, after DB initialization, in
// order to start up language servers, etc.
func StartLangServers(ctx context.Context) {
	if _, ok := langservers.CanManage(); !ok {
		return
	}

	startup := true
	conf.Watch(func() {
		defer func() {
			startup = false
		}()
		for _, language := range langservers.Languages {
			// Start language servers that were previously enabled.
			state, err := langservers.State(language)
			if err != nil {
				log15.Error("failed to get language server state", "language", language, "error", err)
				continue
			}
			if state == langservers.StateEnabled {
				// Start language server now.
				if err := langservers.Start(language); err != nil {
					log15.Error("failed to start language server", "language", language, "error", err)
				}
				continue
			}
			if state == langservers.StateDisabled {
				// Stop the language server if it is running.
				_ = langservers.Stop(language)
				continue
			}

			// Do not run the below code to reflect docker state in the config,
			// except on server startup. Otherwise, we could introduce an
			// infinite loop due to langservers.SetDisabled writing to the
			// config and inherently firing conf.Watch again.
			if !startup {
				continue
			}
			// We didn't start/stop the language server. If it is currently
			// running, this indicates that a server admin did so manually e.g. via
			// `docker run`. It is important that we mark this language as enabled
			// in the site config or else it would show up as "disabled" in the
			// admin UI and we would stop it on server shutdown and never start it
			// again.
			info, err := langservers.Info(language)
			if err != nil {
				log15.Error("failed to get language server info", "language", language, "error", err)
				continue
			}
			if !info.Running() {
				// No container for this language running.
				continue
			}

			// Set disabled=false in the site config.
			if err := langservers.SetDisabled(language, false); err != nil {
				log15.Error("failed to mark running language server as enabled", "language", language, "error", err)
				continue
			}
		}
	})
}