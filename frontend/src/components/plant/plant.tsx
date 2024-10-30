import React, { useEffect, useState } from "react";
import {
    LineChart,
    Line,
    Tooltip,
    ResponsiveContainer,
    XAxis,
    YAxis,
} from "recharts";

export interface graphData {
    val: number;
    timestamp: string;
}

interface Plant {
    name: string;
    value: string;
    time: Date;
    history: graphData[];
}

function PlantComponent(plant: Plant) {
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
        <div key={plant.name} className="grid-item">
            <div className="grid-item-container">
                <h2 className="grid-item-header">{plant.name}</h2>
                <h2 className="grid-item-value">{plant.value}%</h2>
            </div>
            <div className="grid-item-timestamp">
                {timeDifference.days > 0
                    ? timeDifference.days + " days ago"
                    : null}
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
            <ResponsiveContainer width="100%" height={75}>
                <LineChart data={plant.history}>
                    {/* <Tooltip /> */}
                    {/* <XAxis dataKey="Timestamp" /> */}
                    {/* <YAxis /> */}
                    <Line
                        type="monotone"
                        dataKey="Val"
                        stroke="#000000"
                        dot={false}
                    />
                </LineChart>
            </ResponsiveContainer>
        </div>
    );
}

export default PlantComponent;
