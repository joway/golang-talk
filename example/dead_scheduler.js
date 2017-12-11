const sleep = async (ms) => {
    return new Promise(resolve => setTimeout(resolve, ms))
}

const dead = async () => {
    while (true) {
        // await sleep(1000)
    }
}

function main() {
    dead()

    while (true) {
        console.log('main')        
    }
}

main()
