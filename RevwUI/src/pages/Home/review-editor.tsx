import { Divider, Note, Tag, Tooltip, useToasts } from "@geist-ui/core";
import { useEffect, useState } from "preact/hooks";
import { revwAPIService } from "../../services/revw-api.service";

export function ReviewEditor({ reviewID }) {
    const { setToast } = useToasts()
    const [isIdValid, setIsValidID] = useState(false)
    const [readerID, setReaderID] = useState("");

    useEffect(() => {
        async function checkIdValid() {
            let revwService = new revwAPIService();
            let reviewId = await revwService.IsReviewIDAvailable(reviewID);
            setIsValidID(reviewId);
        }

        checkIdValid()
    }, [reviewID]);

    const startSession = async ()=>{
        
        let revwService = new revwAPIService();
        let readId = await revwService.BeginRead(reviewID);
        if(typeof readId === "string")
        {
            setToast({text: `could not begin session: '${readId}'`, type: "error"})
            setReaderID("");
        }
        else
        {
            setReaderID(readId.ID);
        }
    }

    return <>
        {!isIdValid ? <Tag type="error">Review with ID '{reviewID}' was not found.</Tag> : <>
        <Tooltip text={'Click to start session.'}><Tag onClick={startSession} style={{cursor: "pointer"}} type="success">Review with ID '{reviewID}' is valid 👆</Tag></Tooltip>
            {readerID == "" ? null : <>
            <Divider style={{marginTop: "20px"}} width={"100%"}></Divider>
            <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                    <Note label={false} type="success">Reader Info</Note>
                    <Note label="Reader">{readerID}</Note>
                    <Note label="Session">{reviewID}</Note>
                    <Note label="Session Created">100101</Note>
            </div>
            </>}
        </>}
    </>
}