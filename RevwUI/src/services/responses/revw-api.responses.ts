
export type RequestReviewResponse = {
    ID: string
}

export type ReviewSessionResponse = {
    ID: string
    Readers: any[]
    CreationDateTime: string
    ReviewText: string
}