// @flow

const context: {
	appURL?: string;
	authorization?: string;
	cacheControl?: string;
	currentUser?: Object;
	userEmail?: string;
	currentSpanID?: string;
	userAgentIsBot?: boolean;
	assetsRoot?: string;
	buildVars?: Object;
	features?: {[key: string]: any};
	hasLinkedGitHub?: boolean;
} = {};

// Sets the values of the context given a JSContext object from the server.
export function reset(ctx: typeof context) {
	Object.assign(context, ctx);
}

export default context;
