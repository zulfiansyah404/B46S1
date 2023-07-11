class Testimonial {
    #quote = ""
    #image = ""

    constructor(quote, image) {
        this.#quote = quote
        this.#image = image
    }

    get quote() {
        return this.#quote
    }

    get image() {
        return this.#image
    }

    get user() {
        throw new Error('there is must be user to make testimonials')
    }

    get testimonialHTML() {
        // apabila karakter pertama this.user = 1 maka itu adalah personal
        // apabila karakter pertama this.user = 2 maka itu adalah company

        if (this.user[0] == "1") {
            return `<div class="testimonial">
                <div class="logo-test">
                    
                    <i class="fas fa-user"></i>
                    
                </div>
                <div class="content-test">
                    <img src="${this.image}">
                    <div class="comment">
                        <p>
                            "${this.quote}"
                        </p>
                    </div>
                    <h2>
                        - ${this.user.substring(1)}
                    </h2>
                </div>
            
            </div>
            `
        } else if (this.user[0] == "2") {
            return `<div class="testimonial">
                <div class="logo-test">
                    
                    <i class="fas fa-building"></i>
                    
                </div>
                <div class="content-test">
                    <img src="${this.image}">
                    <div class="comment">
                        <p>
                            "${this.quote}"
                        </p>
                    </div>
                    <h2>
                        - ${this.user.substring(1)} Company
                    </h2>
                </div>
            
            </div>
            `
        }
    }
}

class UserTestimonial extends Testimonial {
    #user = ""

    constructor(user, quote, image) {
        super(quote, image)
        this.#user = user
    }

    get user() {
        return "1" + this.#user
    }
}

class CompanyTestimonial extends Testimonial {
    #company = ""

    constructor(company, quote, image) {
        super(quote, image)
        this.#company = company
    }

    get user() {
        return "2" + this.#company
    }
}

const testimonial1 = new UserTestimonial("Surya Elidanto", "GG gaming", "https://images.unsplash.com/photo-1580477667995-2b94f01c9516?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80")

const testimonial2 = new UserTestimonial("Guswandi", "Keren kamu bang", "https://images.unsplash.com/photo-1541562232579-512a21360020?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

const testimonial3 = new CompanyTestimonial("Dumbways", "Apasih ga jelas", "https://images.unsplash.com/photo-1578632767115-351597cf2477?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

// const testimonial4 = new Testimonial("Apasih ga jelas", "https://images.unsplash.com/photo-1578632767115-351597cf2477?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

// console.log(testimonial1, testimonial2, testimonial3)

let testimonialData = [testimonial1, testimonial2, testimonial3]

let testimonialHTML = ""

for (let i = 0; i < testimonialData.length; i++) {
    testimonialHTML += testimonialData[i].testimonialHTML
}

document.getElementById("testimonials").innerHTML = testimonialHTML