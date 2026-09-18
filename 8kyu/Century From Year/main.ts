export const centuryFromYear = (year: number): number => {

    let end: number = year % 100

    if (end === 0) {
        return year
    }
    
    return Math.floor(year / 100 + 1) * 100;
};
