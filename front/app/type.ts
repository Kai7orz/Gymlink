export type ExerciseRecordType = {
    id: number,
    user_id: number,
    user_name: string,
    exercise_image: string,
    exercise_time: number,
    exercise_date: string, // 簡易のために文字列にしているがソートするなら Date 型にした方がいい
    comment: string,
    likes_count: number,
}

export type Illustration = {
    src: string
}