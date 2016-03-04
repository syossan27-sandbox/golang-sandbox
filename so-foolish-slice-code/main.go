numbers := []int{1, 2, 3, 4, 5}

log(numbers)         // 1. [1 2 3 4 5]
log(numbers[2:])     // 2. [3 4 5]
log(numbers[1:3])    // 3. [2 3]

// 注意：負のインデックスは使用できません

// Pythonのようにnumbers[:-1]というコードは動かないので、
// 負のインデックスを使いたい場合は以下のようにします。
// it's so ugly, :-(
log(numbers[:len(numbers)-1])    // 4. [1 2 3 4]

// 6を追加したい場合
numbers = append(numbers, 6)

log(numbers) // 5. [1 2 3 4 5 6]

// 3を削除したい場合
numbers = append(numbers[:2], numbers[3:]...)

log(numbers)    // 6. [1 2 4 5 6]

// 幾つかの数字を追加したい？心配しなくても大丈夫。
// Go言語には"共通の"ベストプラクティスがあるからね！
// oh...Golang developer is so foolish.
numbers = append(numbers[:2], append([]int{3}, numbers[2:]...)...)

log(numbers)    // 7. [1 2 3 4 5 6]

// スライスをコピーするにはこうしよう
copiedNumbers := make([]int, len(numbers))
copy(copiedNumbers, numbers)

// 8. [1 2 3 4 5 6]
log(copiedNumbers)

// まだまだあるよ！
