function correctTail(body, tail) {
    let end = body.at(-1)
    
    if (end === tail) {
        return true
    } else {
        return false
    }
}