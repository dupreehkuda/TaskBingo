export async function _Submit(event: any): Promise<number> {
    event.preventDefault()
    const data = new FormData(event.target);

    const name = data.get('name') as string;
    const task1 = data.get('task1') as string;
    const task2 = data.get('task2') as string;
    const task3 = data.get('task3') as string;
    const task4 = data.get('task4') as string;
    const task5 = data.get('task5') as string;
    const task6 = data.get('task6') as string;
    const task7 = data.get('task7') as string;
    const task8 = data.get('task8') as string;
    const task9 = data.get('task9') as string;
    const task10 = data.get('task10') as string;
    const task11 = data.get('task11') as string;
    const task12 = data.get('task12') as string;
    const task13 = data.get('task13') as string;
    const task14 = data.get('task14') as string;
    const task15 = data.get('task15') as string;
    const task16 = data.get('task16') as string;


    const newResp = {
        id: "",
        pack: {
            title: name,
            tasks: [task1, task2, task3, task4, 
                task5, task6, task7, task8, 
                task9, task10, task11, task12, 
                task13, task14, task15, task16]
        }
    }

    const res = await fetch('https://taskbingo.com/api/task/setTaskPack', {
        method: 'POST',
        headers: {'Origin': 'taskbingo.com'},
        body: JSON.stringify(newResp)
    })

    return res.status
}