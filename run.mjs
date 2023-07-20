const mArgv = argv;

function main() {
    const build = mArgv.build || mArgv.b;
    let run = false;
    if (build) {
        run = true;
        within(async () => {
            cd(build);
            $`goreleaser build --snapshot --clean`;
        });
    }
    if (!run) {
        echo("do nothing");
    }
}

process.nextTick(main);
