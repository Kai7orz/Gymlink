export type exerciseRecordType = {
    id: number,
    imageUrl: string,
    time: number,
    date: string, // 簡易のために文字列にしているがソートするなら Date 型にした方がいい
    comment: string,
}