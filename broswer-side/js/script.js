const form = document.getElementById("form");
const preRespone = document.getElementById("respone");

form.addEventListener("submit", function(e){
    // e = event (browser context)
    e.preventDefault();
    // ^ You say, Ok browser don't submit, I wil handle this.
    console.log(e.target);
    Call_API(e.target);
})

function Call_API(form_){
    fetch("http://localhost:3001/", {
    method: 'POST',  
    body: new FormData(form_)
})
.then((rep)=> rep.text())
.then((txt)=>{
    preRespone.innerHTML = txt;
})
;
}