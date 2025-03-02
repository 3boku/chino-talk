const postMessage = async () => {
    try {
        const input = document.querySelector(".input");
        const message = input.value;
        const feild = document.querySelector(".text3");

        if (message.trim() === "") {
            alert("Please enter a message");
            return;
        }

        feild.innerHTML = "치노가 생각 중 입니다.";

        input.value = "";

        // Send the message to the server
        const result = await fetch("http://localhost:8080/chat", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },

            body: JSON.stringify({ message }),
        }).then((r) => r.json());

        console.log(result);

        feild.innerHTML = result.message;
    } catch (error) {
        console.error(error);
        feild.innerHTML = "죄송해요.. 지금은 말할 기분이 아니에요..";
    }
};
