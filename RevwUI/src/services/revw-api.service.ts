import { RequestReviewResponse } from "./responses/revw-api.responses";

export class revwAPIService {
    private url = "/api/v1/";

    public async Alive() : Promise<boolean>
    {
        const callUrl = this.url+"alive"

        try {
            const result = await fetch(callUrl)
            const json = await result.json()
            if(json["Status"] === undefined || json["Status"] !== "alive")
            {
                return false;
            }
            return true;
        } catch(e)
        {
            console.error("could not fetch "+callUrl);
            console.error(e)
            return false;
        }
    }

    public async RequestReview(text: string) : Promise<RequestReviewResponse | string>
    {
        const callUrl = this.url+"review/request/begin"

        const data = {
            "Text": text,
        };

        try
        {
            const result = await fetch(callUrl, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })

            if(result.status != 200)
            {
                // Error string
                const decoder = new TextDecoder();
                const decodedString = decoder.decode((await result.body.getReader().read()).value);
                return decodedString;
            }

            
            const resultJson = await result.json()
            if(resultJson["ID"] == undefined)
            {
                console.error("Got invalid json from request on URL "+callUrl);
                console.error(resultJson)
                return undefined;
            }

            return resultJson as RequestReviewResponse;
        } catch(e)
        {
            console.error("could not fetch "+callUrl);
            console.error(e)
            return e;
        }
    }

    public async IsReviewIDAvailable(reviewId: string) : Promise<boolean>
    {
        const callUrl = this.url+"review/get"

        const data = {
            "ReviewID": reviewId,
        };

        try
        {
            const result = await fetch(callUrl, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })

            if(result.status != 200)
            {
                // Error string
                const decoder = new TextDecoder();
                const decodedString = decoder.decode((await result.body.getReader().read()).value);
                console.error(decodedString)
                return false;
            }

            
            const resultJson = await result.json() 
            if(resultJson)
            {
                return true;
            }
            return false;
        } catch(e)
        {
            console.error("could not fetch "+callUrl);
            console.error(e)
            return false;
        }
    }

    public async BeginRead(reviewID: string) : Promise<RequestReviewResponse | string>
    {
        const callUrl = this.url+"review/read/begin"

        const data = {
            "ID": reviewID,
        };

        try
        {
            const result = await fetch(callUrl, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })

            if(result.status != 200)
            {
                // Error string
                const decoder = new TextDecoder();
                const decodedString = decoder.decode((await result.body.getReader().read()).value);
                return decodedString;
            }

            
            const resultJson = await result.json()
            if(resultJson["ID"] == undefined)
            {
                console.error("Got invalid json from request on URL "+callUrl);
                console.error(resultJson)
                return undefined;
            }

            return resultJson as RequestReviewResponse;
        } catch(e)
        {
            console.error("could not fetch "+callUrl);
            console.error(e)
            return e;
        }
    }

    public async GetComment()
    {

    }

    public async AddComment()
    {

    }

    public async DeleteComment()
    {

    }

    public async EditComment()
    {

    }
}