let menu = document.querySelector('#menu-btn');
let navbar = document.querySelector('.navbar');

menu.onclick = () =>{
    menu.classList.toggle('fa-times');
    navbar.classList.toggle('active');
}

window.onscroll = () =>{
    menu.classList.remove('fa-times');
    navbar.classList.remove('active');
}

document.querySelector('.book-appointment-button').addEventListener('click',(event)=>{
    var messageElement = document.querySelector(".book-appointment-button");

    messageElement.innerHTML = "Please  Sign In to Book an Appointment";

    const id = setTimeout(()=>{
        var messageElement = document.querySelector(".book-appointment-button");
        messageElement.innerHTML = "Book Appoinment";
        event.preventDefault()
    },3500)
})





function signOut() {
    // Clear session storage
    sessionStorage.clear();
    // Redirect to the sign-in page (replace 'signin.html' with the actual sign-in page URL)
    window.location.href = '/signin';
}
// Attach the sign-out function to the button click event
document.getElementById('signout-button').addEventListener('click', signOut);










function DisplayBlogs(){

    fetch("/listblogs", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(),
      })
        .then(response => response.json())
        .then(data => {
            if(data.message){
                let html = ""
                html += `<h1 class="heading">our <span>blogs</span></h1>
                <div class="box-container">`
                data.message.forEach((element,index) => {
                    if (index > 11){
                        return
                    }
                    html += `<div class="box">
                            <div class="image">
                                <img src="${element.pictures}" alt="">
                            </div>
                            <div class="content">
                                <div class="icon">
                                    <a href="#"><i class="fas fa-calendar"></i>${element.createdtime}</a>
                                    <a href="#"><i class="fas fa-user"></i> by ${element.author}</a>
                                </div>
                                <h3>${element.title}</h3>
                                <p>${element.content}</p>
                                <a href="#" class="btn">learn more <span class="fas fa-chevron-right"></span> </a>
                            </div>
                        </div>`  
                });
                html += "</div>";
                document.getElementById("blogs").innerHTML = html;
            }
        })
        .catch(error => {
          showToast(error.message, "danger", 5000);
        });
    
}
DisplayBlogs()

    


function DispalyCounts(){

    fetch("/listcounts", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(),
      })
        .then(response => response.json())
        .then(data => {
            if(data.message){
               document.getElementById("doctorcount").innerHTML = data.message.doctorcount + "+"
               document.getElementById("patientcount").innerHTML  = data.message.patientcount + "+"
            }
        })
        .catch(error => {
          showToast(error.message, "danger", 5000);
        });
    
}
DispalyCounts()




    

