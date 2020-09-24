function bsearch (a ,n){
    var top = a.length-1
    var bot = 0

    while(bot<=top){
        var mid = Math.floor((top+bot)/2)
        if (n === a[mid]) return mid 
        if (n>a[mid]) bot = mid + 1
        if (n<a[mid]) top = mid - 1
    }
    return null
}

var t = bsearch([1, 3, 4, 6, 7, 8], 4)
console.log('t=', t)
var t = bsearch([1, 3, 4, 6, 7, 8], 5)
console.log('t=', t)
var t = bsearch([1, 3, 4, 6, 7, 8], 6)
console.log('t=', t)
var t = bsearch([1, 3, 4, 6, 7, 8], 9)
console.log('t=', t)