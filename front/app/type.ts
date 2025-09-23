export type exerciseRecordType = {
    id: number,
    image: string,
    time: number,
    date: string, // 簡易のために文字列にしているがソートするなら Date 型にした方がいい
    comment: string,
}