function createPhoneNumber(numbers) {
    let num = numbers.join("");

    return `(${num.slice(0, 3)}) ${num.slice(3, 6)}-${num.slice(6)}`;
}