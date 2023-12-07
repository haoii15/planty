import React, { useEffect, useState } from "react";

interface Plant {
    name: string;
    value: string;
    time: Date;
}

function Plant(plant: Plant) {
    const [timeDifference, setTimeDifference] = useState({
        days: 0,
        hours: 0,
        minutes: 0,
        seconds: 0,
    });

    const calculateTimeDifference = () => {
        const now = new Date();
        const then = new Date(plant.time);
        const difference = now.getTime() - then.getTime();

        if (difference > 0) {
            const days = Math.floor(difference / (1000 * 60 * 60 * 24));
            const hours = Math.floor(
                (difference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
            );
            const minutes = Math.floor(
                (difference % (1000 * 60 * 60)) / (1000 * 60)
            );
            const seconds = Math.floor((difference % (1000 * 60)) / 1000);

            setTimeDifference({ days, hours, minutes, seconds });
        } else {
            // If the target date has passed, set the time difference to 0
            setTimeDifference({ days: 0, hours: 0, minutes: 0, seconds: 0 });
        }
    };

    useEffect(() => {
        calculateTimeDifference();
    }, []);

    // calculateTimeDifference();
    return (
        <div key={plant.name} className="box">
            <h2>{plant.name}</h2>
            {plant.value}%
            <br />
            {timeDifference.days > 0 ? timeDifference.days + " days ago" : null}
            {timeDifference.hours > 1 && timeDifference.days === 0
                ? timeDifference.hours + " hours ago"
                : null}
            {timeDifference.hours === 1 && timeDifference.days === 0
                ? timeDifference.hours + " hour ago"
                : null}
            {timeDifference.minutes > 1 &&
            timeDifference.hours === 0 &&
            timeDifference.days === 0
                ? timeDifference.minutes + " minutes ago"
                : null}
            {timeDifference.minutes === 1 &&
            timeDifference.hours === 0 &&
            timeDifference.days === 0
                ? timeDifference.minutes + " minute ago"
                : null}
            {timeDifference.minutes === 0 &&
            timeDifference.hours === 0 &&
            timeDifference.days === 0
                ? " >1 minute ago"
                : null}
        </div>
    );
}

export default Plant;
