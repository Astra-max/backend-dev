const BASE_URL = "http://localhost:8080"

let userInputs = {FileName: "", Text: ""}

const postRunOne = async () => {
    try {
        const res = await fetch(`${BASE_URL}/ascii-art`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(userInput)
        })

        if (!res.ok) {
            const data = await res.json()
            console.log(data)
        }

    } catch (error) {
        console.table(error)
    }    
}

const postRunAll = async () => {
    const results = validateInput(userInputs)
    try {
        const res = await fetch(`${BASE_URL}/ascii-art`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(results)
        })

        if (res.ok) {
            const data = await res.json()
            console.log(data)
        }
    } catch (error) {
        console.table(error)
    }
}

const validateInput = (data) => {
    const {Text, FileName} = data
    if (Text.length === 0 || Text === "") {
        return "Missing text input"
    }

    if (FileName.length === 0 || FileName === "") {
        return "Missing a banner file"
    }

    return data
}

const getUserInputs = (event) => {
    const {name, value} = event.target
    userInputs[name] = value
}

const errMessage = (err) => {
    const createP = document.createElement('p')
    const motherTag = document.querySelector(".console-out")
    createP.className = "error-p"
    createP.innerText = err
    createP.style.color = "red"
    createP.style.padding = "2rem 0rem"
    motherTag.appendChild(createP)
}