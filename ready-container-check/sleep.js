const fs = require('fs');

const run = async () => {
    for (let i = 0; i < 2 * 60 * 24 * 30 * 365; i++) {
        console.log("Sleeping for 30 seconds");
        await new Promise(r => setTimeout(r, 30000));

        if (i === 2) {
            fs.writeFileSync('./2min', 'ready');
        }
        if (i === 2 * 5) {
            fs.writeFileSync('./5min', 'ready');
        }
        if (i === 2 * 10) {
            fs.writeFileSync('./10min', 'ready');
        }
    }
}

run();