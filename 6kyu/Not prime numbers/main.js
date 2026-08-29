function notPrimes(a,b){
    let results = [];

    for (let n = a; n < b; n++) {
        let digits = String(n);
        let validDigits = true;

        for (let digit of digits) {
            if (!"2357".includes(digit)) {
                validDigits = false
                break;
            }
        }

        if (!validDigits) continue

        let isPrime = true;

        if (n < 2) {
            isPrime = false
        }

        for (let i = 2; i * i <= n; i++) {
            if (n % i === 0) {
                isPrime = false;
                break;
            }
        }

        if (!isPrime) {
            results.push(n);
        }
    }

    return results
}

