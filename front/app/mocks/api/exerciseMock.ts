// いいね機能作るなら record id も必要かも
const exerciseMock = {
        id: 1,
        image: '/images/sportImage.png',
        time: 45,
        date: '2025-09-20',
        comment: 'いい運動',
    };

const exerciseMocks = [
    {
        id: 1,
        image: '/images/sportImage.png',
        time: 45,
        date: '2025-09-20',
        comment: 'いい運動',

    },
    {
        id: 2,
        image: '/images/sportImage.png',
        time: 45,
        date: '2025-09-20',
        comment: 'いい運動',

    },
    {
        id: 3,
        image: '/images/sportImage.png',
        time: 45,
        date: '2025-09-20',
        comment: 'いい運動',

    },
        {
        id: 4,
        image: '/images/sportImage.png',
        time: 45,
        date: '2025-09-20',
        comment: 'いい運動',

    },
    {
        id: 5,
        image: '/images/sportImage.png',
        time: 45,
        date: '2025-09-20',
        comment: 'いい運動',

    },
    {
        id: 6,
        image: '/images/sportImage.png',
        time: 45,
        date: '2025-09-20',
        comment: 'いい運動',

    }
]

if(exerciseMocks[0]) {
exerciseMocks[0].date = '2025-09-18';
exerciseMocks[0].comment = '走れた'
}

if(exerciseMocks[1]) {
exerciseMocks[1].date = '2025-09-19';
exerciseMocks[1].comment = 'いい感じに走れた'
}


if(exerciseMocks[2]) {
exerciseMocks[2].date = '2025-09-20';
exerciseMocks[2].comment = '疲れた'
}


export { exerciseMock,exerciseMocks };