var fs = require('fs');

data = fs.readFileSync("./responses.txt", 'utf-8').slice(0,-1);

lines = data.split("\n");
console.log("fname,gender");
for(var i=0;i<lines.length;i++){
  var line = lines[i];
  line = JSON.parse(line);
  if(Math.abs(line.scale) >= 0.5){
    console.log(line.firstName + "," + line.gender);
  }
}

