function l(a,v){try{a.splice(li(a,v)+1,0,v);}catch(e){}}
function m(a,v){try{a.splice(mi(a,v)+1,0,v);}catch(e){}}

function li(a,v,i) {
  i=i||a.length - 1;
  if(a.length==0||i==0){
    return 0;
  } else if(a[i]==v){
    throw'e';
  } else if(a[i]<v){
    return i;
  } else {
    return li(a,v,i-1);
  }
}

function mi(a, v, i) {
  i = Math.floor(a.length/2);
  
  if (a.length==0||i<=0) {
    return 0;
  
  } else if(a[i] == v) {
    throw'e';
  
  } else if(a[i]<v && i>=a.length-1) {
    return i;
  
  } else if(a[i]<v && v<a[i+1]) {
    return i;
  
  } else if(a[i]<v) {
    return mi(a.slice(i, a.length - 1),v);  
  
  } else if(a[i]>v) {
    return mi(a.slice(0, i),v);
  }
}

function dblLinear(n) {
  var start = Date.now();
  var u = [1];
  var s = Math.ceil(n*0.65);
  for (var i = 0; i < s; i++) {
    var x = u[i];
    m(u, 2 * x + 1);
    l(u, 3 * x + 1);
  }
  console.log(Date.now() - start);
  return u[n];
}
