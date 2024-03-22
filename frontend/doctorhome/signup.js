
function Submitform(formData) {

  console.log("Request sent");
  fetch("/signupdoctor", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(formData),
  })
    .then(response => response.json())
    .then(data => {
      if (data && data.value) {
        showToast(data.value, "success", 5000);

      } else {
        showToast(data.error, "danger", 5000);
      }

      WindowWait('/signin')
        document.getElementById("name").value = '',
        document.getElementById("email").value = '',
        document.getElementById("password").value = '',
        document.getElementById("confirm_password").value = ''
    })
    .catch(error => {
      showToast(data.error, "danger", 5000);
    });
};



let icon = {
  success:
    '<span class="material-symbols-outlined">task_alt</span>',
  danger:
    '<span class="material-symbols-outlined">error</span>',
  warning:
    '<span class="material-symbols-outlined">warning</span>',
  info:
    '<span class="material-symbols-outlined">info</span>',
};

const showToast = (
  message = "Sample Message",
  toastType = "info",
  duration = 5000) => {
  if (
    !Object.keys(icon).includes(toastType))
    toastType = "info";

  let box = document.createElement("div");
  box.classList.add(
    "toast", `toast-${toastType}`);
  box.innerHTML = ` <div class="toast-content-wrapper"> 
                    <div class="toast-message">${message}</div> 
                    <div class="toast-progress"></div> 
                    </div>`;
  duration = duration || 5000;
  box.querySelector(".toast-progress").style.animationDuration =
    `${duration / 1000}s`;

  let toastAlready =
    document.body.querySelector(".toast");
  if (toastAlready) {
    toastAlready.remove();
  }

  document.body.appendChild(box)
};

function WindowWait(str) {
  setTimeout(() => {
    window.location.href = str
  }, 3000)

}




// Attach the sign-out function to the button click event
document.getElementById('signout-button').addEventListener('click', ()=>{
  console.log("In Signout")
  sessionStorage.clear();
  // Redirect to the sign-in page (replace 'signin.html' with the actual sign-in page URL)
  window.location.href = '/home';
});






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