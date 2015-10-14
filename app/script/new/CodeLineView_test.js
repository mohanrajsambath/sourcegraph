import sandbox from "../testSandbox";

import React from "react";

import CodeLineView from "./CodeLineView";

describe("CodeLineView", () => {
	it("should render tokens", () => {
		sandbox.renderAndExpect(<CodeLineView lineNumber={42} tokens={[{Label: "foo", Class: "a"}, {Label: "bar"}, {Label: "baz", Class: "c"}]} />).to.eql(
			<tr className="line">
				<td className="line-number" data-line={42}></td>
				<td className="line-content">
					<span className={"a"} key={0}>{"foo"}</span>
					<span className={undefined} key={1}>{"bar"}</span>
					<span className={"c"} key={2}>{"baz"}</span>
				</td>
			</tr>
		);
	});
});
