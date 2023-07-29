

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

    const form = document.getElementById("form-project")
    const formData = new FormData(form)

    sendDataToServer(formData)
}   

function sendDataToServer(formData) {
    const url = "http://localhost:8000/"

    fetch(url, {
        method: "POST",
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            console.log(data)
            window.location.href = "http://localhost:8000/"
        })
        .catch(error => {
            return console.log(error)
        })
}

