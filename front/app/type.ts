export type RecordType = {
    id: number,
    user_id: number,
    user_name: string,
    presigned_image: string,
    clean_up_time: number,
    clean_up_date: string, // 簡易のために文字列にしているがソートするなら Date 型にした方がいい
    comment: string,
    likes_count: number,
}

export type Illustration = {
    src: string
}