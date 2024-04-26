import { Input } from "@geist-ui/core";
import { useState } from "preact/hooks";
import { ReviewEditor } from "./review-editor";


export function ReviewRequest() {
    const [reviewID, setReviewID] = useState("")

    const handleIDChange = (event)=>{
        setReviewID(event.target.value);
    }

    return <>
        <div style={{ display: "flex", flexDirection: "column", alignItems: "center" }}>
            <input style={{marginBottom: "20px"}} placeholder="Review ID" value={reviewID} onChange={handleIDChange}></input>
            {reviewID != "" ? <ReviewEditor reviewID={reviewID}></ReviewEditor> : null}
        </div>
    </>
}