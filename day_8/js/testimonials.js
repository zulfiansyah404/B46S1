// class Testimonial {
//     #quote = ""
//     #image = ""

//     constructor(quote, image) {
//         this.#quote = quote
//         this.#image = image
//     }

//     get quote() {
//         return this.#quote
//     }

//     get image() {
//         return this.#image
//     }

//     get user() {
//         throw new Error('there is must be user to make testimonials')
//     }

//     get testimonialHTML() {
//         // apabila karakter pertama this.user = 1 maka itu adalah personal
//         // apabila karakter pertama this.user = 2 maka itu adalah company

//         if (this.user[0] == "1") {
//             return `<div class="testimonial">
//                 <div class="logo-test">
                    
//                     <i class="fas fa-user"></i>
                    
//                 </div>
//                 <div class="content-test">
//                     <img src="${this.image}">
//                     <div class="comment">
//                         <p>
//                             "${this.quote}"
//                         </p>
//                     </div>
//                     <h2>
//                         - ${this.user.substring(1)}
//                     </h2>
//                 </div>
            
//             </div>
//             `
//         } else if (this.user[0] == "2") {
//             return `<div class="testimonial">
//                 <div class="logo-test">
                    
//                     <i class="fas fa-building"></i>
                    
//                 </div>
//                 <div class="content-test">
//                     <img src="${this.image}">
//                     <div class="comment">
//                         <p>
//                             "${this.quote}"
//                         </p>
//                     </div>
//                     <h2>
//                         - ${this.user.substring(1)}'s Company
//                     </h2>
//                 </div>
            
//             </div>
//             `
//         }
//     }
// }

// class UserTestimonial extends Testimonial {
//     #user = ""

//     constructor(user, quote, image) {
//         super(quote, image)
//         this.#user = user
//     }

//     get user() {
//         return "1" + this.#user
//     }
// }

// class CompanyTestimonial extends Testimonial {
//     #company = ""

//     constructor(company, quote, image) {
//         super(quote, image)
//         this.#company = company
//     }

//     get user() {
//         return "2" + this.#company
//     }
// }

// const testimonial1 = new UserTestimonial("Jotaro Kujo", "Yare Yare, mantab banget!", "img/jotaro.png")

// const testimonial2 = new UserTestimonial("Ruby Hoshino", "Mantab banget oni-chan", "https://japan.postsen.com/content/uploads/2023/04/18/3deb9d4074.jpg")

// const testimonial3 = new CompanyTestimonial("Umbrella", "Boleh juga, bagi dulu virus Lorem ipsum dolor sit amet consectetur adipisicing elit. Adipisci deleniti incidunt magni assumenda, quos voluptas enim saepe dolorum nulla amet, sed dolorem odit, doloribus itaque molestiae molestias exercitationem. Vero, deserunt.", "https://e7.pngegg.com/pngimages/484/820/png-clipart-umbrella-corps-umbrella-corporation-logo-corporation-s-umbrella-triangle.png")

// // const testimonial4 = new Testimonial("Apasih ga jelas", "https://images.unsplash.com/photo-1578632767115-351597cf2477?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

// // console.log(testimonial1, testimonial2, testimonial3)

// let testimonialData = [testimonial1, testimonial2, testimonial3]

// let testimonialHTML = ""

// for (let i = 0; i < testimonialData.length; i++) {
//     testimonialHTML += testimonialData[i].testimonialHTML
// }

// document.getElementById("testimonials").innerHTML = testimonialHTML

// testimonialsHTML(){
//         
// DATA
const testiData = [
    {
        image : "img/jotaro.png",
        text : "Yare Yare, mantab banget!",
        autors : "Jotaro Kujo",
        rating : 5,
        isCompany : false
    },
    {
        image : "https://japan.postsen.com/content/uploads/2023/04/18/3deb9d4074.jpg",
        text : "Mantab banget oni-chan",
        autors : "Ruby Hoshino",
        rating : 4,
        isCompany : false
    },
    {
        image : "https://e7.pngegg.com/pngimages/484/820/png-clipart-umbrella-corps-umbrella-corporation-logo-corporation-s-umbrella-triangle.png",
        text : "Boleh juga, bagi dulu virus Lorem ipsum dolor sit amet consectetur adipisicing elit. Adipisci deleniti incidunt magni assumenda, quos voluptas enim saepe dolorum nulla amet, sed dolorem odit, doloribus itaque molestiae molestias exercitationem. Vero, deserunt.",
        autors : "Umbrella",
        rating : 3,
        isCompany : true
    },
    {
        image : "https://asset.kompas.com/crops/9XCiAwPgGDlJT3H_vUhGQIqs0-A=/0x76:1280x929/750x500/data/photo/2021/12/10/61b2c12de1994.jpeg",
        text : "Boleh sih, tapi jangan lupa tanam kopi dulu",
        autors : "Herman William Daendles",
        rating : 1,
        isCompany: true
    }
    ,
    {
        image : "https://cdn.oneesports.id/cdn-data/wp-content/uploads/sites/2/2021/01/Valorant_YoruLaunchImage.jpg",
        text : "Not Bad lah",
        autors : "Yoru",
        rating : 2,
        isCompany: false
    }
]
    
function testimonialsAll() {
    let testimonialsHTML = ""

    testiData.forEach((items) => { 
        let icon = ""
        if (items.isCompany) {
            icon = `<i class="fas fa-building"></i>`
        } else {
            icon = `<i class="fas fa-user"></i>`
        }

        testimonialsHTML += `
        <div class="testimonial">
            <div class="logo-test">
                ${icon}   
            </div>
            <div class="content-test">
                <img src="${items.image}" alt="">
                <div class="rating"> 
            `
        
        for (let i = 0; i < items.rating; i++) {
            testimonialsHTML += `
                <i class="fa-solid fa-star"></i>
            `
        }

        testimonialsHTML += `
                </div>
                <div class="comment">
                    <p>
                        "${items.text}"
                    </p>
                </div>
                <h2>
                    - ${items.autors}
                </h2>
                
            </div>
        
        </div>
        `
    })     
    document.getElementById("testimonials").innerHTML = testimonialsHTML
}

testimonialsAll()

// FILTER
function FilterTestimonial(rate) {
    let filteredTestimonialHTML = ""

    const filteredData = testiData.filter((items) => {
        return items.rating === rate
    })

    testiData.forEach((items) => { 
        let icon = ""
        if (items.isCompany) {
            icon = `<i class="fas fa-building"></i>`
        } else {
            icon = `<i class="fas fa-user"></i>`
        }

        testimonialsHTML += `
        <div class="testimonial">
            <div class="logo-test">
                ${icon}   
            </div>
            <div class="content-test">
                <img src="${items.image}" alt="">
                <div class="rating"> 
            `
        
        for (let i = 0; i < items.rating; i++) {
            testimonialsHTML += `
                <i class="fa-solid fa-star"></i>
            `
        }

        testimonialsHTML += `
                </div>
                <div class="comment">
                    <p>
                        "${items.text}"
                    </p>
                </div>
                <h2>
                    - ${items.autors}
                </h2>
                
            </div>
        
        </div>
        `
    })     
    document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
}