const postData = (event) => {
    let [username, password] = getUserData()
    try {
        const sendData = async () => {
            const res = await fetch("http://localhost:8080", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({username, password})
            })
            if (res.ok) {
                successMessage()
                alert("sent succefully")
            }
        }
        sendData()
    } catch (error) {
        logError(error)
    }
}

const getUserData = () => {
    const userName = document.querySelector(".username").value
    const password = document.querySelector(".password").value

    return [userName, password]
}

const successMessage = () => {
    const btn = document.querySelector(".s-btn")
    btn.innerHTML = "submitted successfully"
    btn.style.backgroundColor = "green"
    btn.style.cursor = "not-allowed"
    btn.style.opacity = "0.7"
    btn.disabled = true
}

const logError = (msg) => {
    const createP = document.createElement('p')
    createP.textContent = msg
    createP.style.color = "red"
    const section = document.querySelector('.user-data')

    if (!section) return
    section.appendChild(createP)
}