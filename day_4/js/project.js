let firstProject = {
    name : "Project Tubes III STIMA",
    duration : "1 bulan",
    description : "Project Tubes III STIMA Teknik Informatika ITB merupakan project yang dikerjakan oleh 3 orang mahasiswa Teknik Informatika ITB. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma.",
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
        return `${centuries} abad`
    } else if (years > 0) {
        return `${years} tahun`
    } else if (months > 0) {
        return `${months} bulan`
    } else {
        return `${days} hari`
    }
}

const addProject = (event) => {
    event.preventDefault()

    let name = document.getElementById('project-name').value
    let description = document.getElementById('description').value
    let start_date = document.getElementById('start-date').value
    let end_date = document.getElementById('end-date').value

    // Cek apakah name, description, start_date, end_date kosong atau tidak
    if (name == '') {
        alert('Nama Project tidak boleh kosong')
        return
    }
    if (start_date == '') {
        alert('Tanggal Mulai tidak boleh kosong')
        return
    }
    if (end_date == '') {
        alert('Tanggal Selesai tidak boleh kosong')
        return
    }
    if (description == '') {
        alert('Deskripsi Project tidak boleh kosong')
        return
    }

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
    if (document.getElementById('golang').checked) {
        listIconTech.push(golang_icon)
    }

    //  Cek apakah listIconTech kosong atau tidak
    if (listIconTech.length == 0) {
        alert('Minimal pilih 1 teknologi')
        return
    }

    //  Cek apakah image kosong atau tidak
    let image = document.getElementById('image').files
    
    if (image.length == 0) {
        alert('Gambar tidak boleh kosong')
        return
    }

    image = URL.createObjectURL(image[0])

    duration = getDuration(start_date, end_date)
    if (duration == false) {
        alert('Tanggal Selesai tidak boleh lebih kecil dari Tanggal Mulai')
        return
    }


    console.log(start_date)
    console.log(typeof start_date)

    let project = {
        name, 
        duration,
        description, 
        listIconTech,
        href : '#',
        image
    }

    listProject.push(project)
    renderProject()
}   

// Fungsi untuk merender list project
const renderProject = () => {
    document.getElementById("list-project").innerHTML = '' // Inisialisasi elemen list Project masih kosong
    for (let i = 0; i < listProject.length; i++) {
        document.getElementById("list-project").innerHTML += `
        <div class="project">
            <div class="header-project-item">
                <div class="circle-red"></div>
                <div class="circle-yellow"></div>
                <div class="circle-green"></div>
            </div>
            <div class="blog-image">
                <img src="${listProject[i].image}" alt="">
            </div>
            <div class="project-items">
                
                <h2><a href="${listProject[i].href}">${listProject[i].name}</a></h2>
                <h3><div class="duration-label">Durasi : </div>${listProject[i].duration}</h3>
                <div class="paragraph">
                    <p>${listProject[i].description}</p>
                </div>
                
                <div class="icons">
                    ${listProject[i].listIconTech.join('')}
                </div>
                <div class="button-project">
                    <button style="background:#fb4c33;"><a>Edit</a></button>
                    <button style="background:rgba(255, 0, 0, 0.804);"><a>Delete</a></button>
                </div>
            </div>
        </div>
        `
    }
}
