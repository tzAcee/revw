import { useState } from "preact/hooks";
import { Tag, useToasts } from "@geist-ui/core";
import { revwAPIService } from "../../services/revw-api.service";

export function RequestReview() {
    const { setToast } = useToasts()
    const [text, setText] = useState("");
    const [status, setStatus] = useState("");
    const [succeed, setSuccess] = useState(false)
    const [lastId, setLastId] = useState("");

    const requestReview = async ()=>{
        let revwService = new revwAPIService();
        let response = await revwService.RequestReview(text);
        if(typeof response === "string")
        {
            setSuccess(false)
            setStatus(response);
        }
        else
        {
            const statusText = `review requested: '${response.ID}' 👆`;
            setLastId(response.ID)
            setStatus(statusText)
            setSuccess(true)
            copyToClipboard(response.ID);
        }
    }

    const handleTextChange = (event) => {
        setText(event.target.value);
    };

    const copyToClipboard = (id) => {
        setTimeout(()=>{
            navigator.clipboard.writeText(id)
            .then(() => {
                setToast({ text: `'${id}' was copied to the clipboard.`, delay: 1000 })
            })
            .catch((error) => {
                setToast({ text: `Failed to copy '${id}' to clipboard:`, type:"error"});
            });
        }, 100)

    };

    const handleTagClick = () =>{
        copyToClipboard(lastId)
    }

    return <>
            <style>
                {`
                    textarea {
                        font-size: 0.8rem;
                        letter-spacing: 1px;
                        padding: 10px;
                        width: 60%;
                        height: 20vh;
                        max-width: 100%;
                        line-height: 1.5;
                        border-radius: 5px;
                        border: 1px solid #ccc;
                        box-shadow: 1px 1px 1px #999;
                        margin-bottom: 15px;
                    }
                `}
            </style>
    <span style={{ display: "flex", flexDirection: "column", alignItems: "center" }}>
        <div style={{ display: "flex", alignItems: "center", marginBottom: "20px" }}>
            <button onClick={requestReview}>Request Review</button>
            {status && <Tag
            style={{ marginLeft: "10px", cursor: succeed ? "pointer" : "default" }}
            type={succeed ? "success" : "error"}
            onClick={succeed ? handleTagClick : null}
            >{status}</Tag>}
        </div>
        <textarea placeholder={"Your text you want to get reviewed..."} onChange={handleTextChange} value={text}></textarea>
    </span>
    </>
}