import { Divider } from "@geist-ui/core";
import { RequestReview } from "./request-review";
import { ReadReview } from "./read-review";

export function Request() {
    return<div style={{ textAlign: "center", marginTop: "30px" }}>
        <RequestReview></RequestReview>
        <Divider>OR</Divider>
        <ReadReview></ReadReview>
    </div>
}