export function findShort(s: string): number {
    const words: string[] = s.split(" ");

    let min: number = words[0].length;

    for (const word of words) {
        if (word.length < min) {
            min = word.length;
        }
    }

    return min;
}