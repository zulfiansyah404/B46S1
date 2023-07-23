let firstProject = {
    name : "String Matching dan Regular Expression dalam Pembuatan ChatGPT Sederhana",
    duration : "1 month",
    description : "Project Tubes III STIMA Teknik Informatika ITB merupakan project yang dikerjakan oleh 3 orang mahasiswa Teknik Informatika ITB. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma.Repository ini berisi aplikasi ChatGPT sederhana berbasis web. Aplikasi ChatGPT sederhana ini dibangun dengan mengimplementasikan algoritma string matching (Knuth-Morris-Pratt dan Boyer-Moore) dan juga regular expression. Aplikasi ini dapat menjawab pertanyaan berdasarkan masukan dari pengguna. Jenis pertanyaan yang dapat dijawab antara lain: Pertanyaan yang terdapat didalam database (menjawab menggunakan algoritma KMP atau BM) Kalkulator sederhana Menanyakan hari berdasarkan tanggal Menambahkan pertanyaan dan jawaban ke dalam database Menghapus pertanyaan dari database Selain itu, history dari pertanyaan yang dimasukkan akan disimpan ke dalam sebuah section pertanyaan untuk sebuah sesi. Pengguna juga dapat menambahkan section baru dengan mengklik add new history. Untuk pembuatan program ini, pembuatan frontend dilakukan dengan menggunakan bahasa pemrograman HTML, CSS, dan juga Javascript dengan menggunakan Framework React. Untuk Backend dari program dibuat dengan menggunakan bahasa pemrograman Golang. Sedangkan untuk penyimpanan database menggunakan MySQL.",
    listIconTech : ['<i class="fa-brands fa-node-js"></i>',
    '<i class="fa-brands fa-react"></i>'],
    image : "img/project.png",
    href : "project-detail.html"
}
let listProject = [firstProject]

// Fungsi untuk menentukan berapa hari, bulan, tahun, atau abad dari start_date dan end_date
// Parameter : start_date, end_date : string(format: yyyy-mm-dd)
const getDuration = (start_date, end_date) => {
    let start = new Date(start_date)
    let end = new Date(end_date)

    let duration = end - start

    // Jika durasi minus, maka return false
    if (duration < 0) {
        return false
    }

    let days = Math.floor((duration) / (1000 * 60 * 60 * 24))
    let months = Math.floor(days / 30)
    let years = Math.floor(months / 12)
    let centuries = Math.floor(years / 100)

    if (centuries > 0) {
        if (centuries > 1) {
            return `${centuries} centuries`
        }
        return `${centuries} century`
    } else if (years > 0) {
        if (years > 1) {
            return `${years} years`
        }
        return `${years} year`
    } else if (months > 0) {
        if (months > 1) {
            return `${months} months`
        }
        return `${months} bulan`
    } else {
        if (days > 1) {
            return `${days} days`
        }
        return `${days} hari`
    }
}


// Prosedur untuk menghapus peringatan inputan kosong
// Parameter : id : string
const removeAlertEmpty = (id) => {
    document.getElementById(id).innerHTML = ''
}

const addProject = (event) => {
    event.preventDefault()

    let name = document.getElementById('name').value
    let description = document.getElementById('description').value
    let start_date = document.getElementById('start-date').value
    let end_date = document.getElementById('end-date').value

    let isDataComplete = true

    // Icon teknologi
    const nodejs_icon = '<i class="fa-brands fa-node-js"></i>';
    const reactjs_icon = '<i class="fa-brands fa-react"></i>';
    const java_icon = '<i class="fa-brands fa-java"></i>';
    const golang_icon = '<i class="fa-brands fa-golang"></i>';

    let listIconTech = []
    if (document.getElementById('node-js').checked) {
        listIconTech.push(nodejs_icon)
    }
    if (document.getElementById('react-js').checked) {
        listIconTech.push(reactjs_icon)
    }
    if (document.getElementById('java').checked) {
        listIconTech.push(java_icon)
    }
    if (document.getElementById('go').checked) {
        listIconTech.push(golang_icon)
    }

    //  Cek apakah listIconTech kosong atau tidak
    if (listIconTech.length == 0) {
        document.getElementById('techHelp').innerHTML = 'Choose at least 1 technology'
        isDataComplete = false
    } else {
        document.getElementById('techHelp').innerHTML = ''
    }

    //  Cek apakah image kosong atau tidak
    let image = document.getElementById('image').files

    image = URL.createObjectURL(image[0])

    duration = getDuration(start_date, end_date)
    if (duration == false) {
        document.getElementById('dateHelp').innerHTML = 'The end date cannot be earlier than the start date'
        return
    }

    // Debug
    // console.log(start_date)
    // console.log(typeof start_date)

    if (!isDataComplete) {
        return
    }

    let project = {
        name, 
        duration,
        description, 
        listIconTech,
        href : 'project-detail.html',
        image
    }

    listProject.push(project)
    renderProject()
}   

// Fungsi untuk merender list project
const renderProject = () => {
    document.getElementById("projects").innerHTML = '' // Inisialisasi elemen list Project masih kosong
    for (let i = 0; i < listProject.length; i++) {
        document.getElementById("projects").innerHTML += `
        <div class="col">
            <div class="project card ms-auto me-auto shadow-lg border-0 border-bottom mb-3"  data-bs-theme="dark">
                <div class="card-header d-flex flex-row-reverse pt-1 pe-3">
                    <!-- font-awesome tempat sampah -->
                    <a href="#" class="text-decoration-none text-danger "><i class="fas fa-trash-alt"></i></a>
                    <!-- font awesome pencil -->
                    <a href="#" class="text-decoration-none text-warning me-3"><i class="fas fa-pencil-alt"></i></a>
                    
                </div>
                <img src="${listProject[i].image}" class="card-img-top mb-3" alt="...">
                <div class="card-body ps-4 pe-4">
                    <h5 class="card-title mb-2"><a href="${listProject[i].href}" class="fw-bold fs-3 text-decoration-none" id="card-title">${listProject[i].name}</a></h5>
                    <div class="card-duration glacial-indifference fs-6 mb-3">
                        <span style="color: #fe7e67;">Duration: </span> ${listProject[i].duration}
                    </div>
                    <p class="card-text fs-7 mb-3">${listProject[i].description}</p>
                    <div class="icon-tech fs-2 text-info-emphasis">
                        ${listProject[i].listIconTech.join(' ')}
                    </div>
                </div>
            </div>
            
            

        </div>
                
        `
    }
}

// Program Utama
renderProject()