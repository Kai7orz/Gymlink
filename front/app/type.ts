export type ExerciseRecordType = {
    id: number,
    user_name: string,
    image_url: string,
    time: number,
    date: string, // 簡易のために文字列にしているがソートするなら Date 型にした方がいい
    comment: string,
    likes_count: number,
}

export type Illustration = {
    src: string
}