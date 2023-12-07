import React, { useEffect, useState } from "react";
import "./styles.css"; // Import your CSS file
import Plant from "./plant/plant";

interface Data {
    name: string;
    value: string;
    time: Date;
    // Add other properties from your API response
}

function MyComponent() {
    const [dataList, setDataList] = useState<Data[]>([]);
    const [refreshCount, setRefreshCount] = useState(0);
    const refreshInterval = 1000; // 5 seconds

    // Define the refresh interval in milliseconds

    useEffect(() => {
        // Function to fetch data and update the state
        const fetchData = async () => {
            try {
                // Replace this with your API endpoint
                // const response = await fetch("http://192.168.68.106:61942/");
                const response = await fetch("http://192.168.68.143:61942/");
                // const response = await fetch("http://192.168.105.209:61942/");
                const data = await response.json();
                setDataList(data);
            } catch (error) {
                console.error("Error fetching data:", error);
            }
        };

        // Fetch data immediately and set up a refresh interval
        fetchData();

        const intervalId = setInterval(() => {
            // Increment the refreshCount to trigger a re-render
            setRefreshCount((prevCount) => prevCount + 1);

            // Fetch data on each refresh
            fetchData();
        }, refreshInterval);

        // Clean up the interval when the component unmounts
        return () => clearInterval(intervalId);
    }, []); // Empty dependency array means this effect runs once on component mount

    return (
        <div className="container">
            <img
                src="/leaf.png"
                alt="failed to load"
                style={{ width: 100, height: 100 }}
            />
            <h1>Planty</h1>
            {dataList.map((item) => (
                <Plant
                    name={item.name}
                    value={item.value}
                    time={item.time}
                ></Plant>
            ))}
        </div>
    );
}

export default MyComponent;
