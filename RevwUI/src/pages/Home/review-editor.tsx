import { Divider, Modal, Note, Tag, Tooltip, useToasts } from "@geist-ui/core";
import { useEffect, useState } from "preact/hooks";
import { revwAPIService } from "../../services/revw-api.service";

type ReaderInfo = {
    sessionID: string
    readerID: string
    sessionCreationTime: Date
    text: string
}

export function ReviewEditor({ reviewID }) {
    const { setToast } = useToasts()
    const [isIdValid, setIsValidID] = useState(false)
    const [readerInfo, setReaderInfo] = useState({sessionID: "", readerID:""} as ReaderInfo)
    const [needCommentModal, setNeedComment] = useState(false)
    const [newCommentIndex, setNewCommentIndex] = useState(-1)
    const [newCommentContent, setNewCommentContent] = useState("");

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
            setReaderInfo({sessionID: "", readerID:""} as ReaderInfo)
        }
        else
        {
            const reviewData = await revwService.GetReviewData(reviewID)
            if(reviewData == undefined)
            {
                setToast({text: `could not get session data of '${reviewID}'`, type: "error"})
                setReaderInfo({sessionID: "", readerID:""} as ReaderInfo)
                return
            }

            const readerInf : ReaderInfo = {
                sessionID: reviewID,
                readerID: readId.ID,
                sessionCreationTime: new Date(reviewData.CreationDateTime),
                text: reviewData.ReviewText
            }
            setReaderInfo(readerInf);
        }
    }

    const onClickEditor = (event)=>{
        if(newCommentIndex != -1)
        {
            return;
        }
        const cursorIndex = event.target.selectionStart; // Get index of cursor position
        setNewCommentIndex(cursorIndex)
        setNeedComment(true)
    }

    const modalCloseHandler = ()=>{
        setNewCommentContent("")
        setNewCommentIndex(-1)
        setNeedComment(false);
    }

    const getCommentLocation = (index) =>{
        const range = 10;

        if(index >= readerInfo.text.length)
        {
            // add | at last pos
            if(readerInfo.text.length < range)
            {
                return readerInfo.text + "|";
            }
            return readerInfo.text.substring(readerInfo.text.length-range) + "|";
        }
        let beforeString = "";
        let afterString = "";

        {
            const stringUntilIndex = readerInfo.text.substring(0, index)
            // add | at last pos
            if(stringUntilIndex.length < range)
            {
                beforeString = stringUntilIndex;
            }
            else
            {
                beforeString = stringUntilIndex.substring(stringUntilIndex.length-range)
            }
        }

        {
            const stringAfterIndex = readerInfo.text.substring(index)
            // add | at last pos
            if(stringAfterIndex.length < range)
            {
                afterString = stringAfterIndex;
            }
            else
            {
                afterString = stringAfterIndex.substring(0, range)
            }
        }

        return "..."+beforeString + " | " + afterString+"...";

    }

    const onCommentContentChange = (event) =>{
        setNewCommentContent(event.target.value)
    }

    const createComment = async ()=>{
        let revwService = new revwAPIService();
        let createCommentResult = await revwService.CreateComment(reviewID, readerInfo.readerID, newCommentContent, newCommentIndex);
        setToast({text: `Create Comment returned with '${createCommentResult}'`})
        setNewCommentContent("")
        setNewCommentIndex(-1)
        setNeedComment(false);
    }

    return <>
        <style>
            {`
                .review-editor {
                    font-size: 0.8rem;
                    letter-spacing: 1px;
                    padding: 10px;
                    width: 75vw;
                    height: 40vh;
                    max-width: 100%;
                    line-height: 1.5;
                    border-radius: 5px;
                    border: 1px solid #ccc;
                    box-shadow: 1px 1px 1px #999;
                    margin-bottom: 15px;
                    margin-top: 20px;
                }
            `}
        </style>

        {!isIdValid ? <Tag type="error">Review with ID '{reviewID}' was not found.</Tag> : <>
        <Tooltip text={'Click to start session.'}><Tag onClick={startSession} style={{cursor: "pointer"}} type="success">Review with ID '{reviewID}' is valid 👆</Tag></Tooltip>
            {readerInfo.readerID == "" ? null : <>
            <Divider style={{marginTop: "20px"}} width={"100%"}></Divider>
            <div style={{ display: "flex", flexDirection: "row", alignItems: "center" }}>
                    <Note label={false} type="success">Reader Info</Note>
                    <Note label="Reader">{readerInfo.readerID}</Note>
                    <Note label="Session">{readerInfo.sessionID}</Note>
                    <Note label="Session Created">{readerInfo.sessionCreationTime.toLocaleString()}</Note>
            </div>
            <div>
                <Note style={{marginTop: "8px", fontSize: "90%"}}>Click the text to add a comment.</Note>
                <textarea onClick={onClickEditor} className={"review-editor"} readOnly value={readerInfo.text}>
                <Modal visible={needCommentModal} onClose={modalCloseHandler}>
                    <Modal.Title>Add new comment</Modal.Title>
                    <Modal.Subtitle>{getCommentLocation(newCommentIndex)}</Modal.Subtitle>
                    <Modal.Content style={{display: "flex", flexDirection: "column"}}>
                    <textarea placeholder={"Comment ..."} onChange={onCommentContentChange} value={newCommentContent}></textarea>
                    <button style={{marginTop:"5px"}}onClick={createComment}>Create</button>
                    </Modal.Content>
                </Modal>

                </textarea>
            </div>
            </>}
        </>}
    </>
}