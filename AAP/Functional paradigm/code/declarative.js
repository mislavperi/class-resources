let list = ["dogs", "hotdogs"]

function upcase(list) {
  let newList = []
  list.forEach(item => {
    newList.push(item.toUpperCase())
  })
  return newList
}

console.log(upcase(list))