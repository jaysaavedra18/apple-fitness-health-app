"use client";

import React from "react";
import {
  LineChart,
  Line,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
} from "recharts";
import ClientOnly from "../utils/clientOnly";

const data = [
  { name: "Page A", uv: 400, pv: 2400, amt: 2400 },
  { name: "Page B", uv: 300, pv: 2000, amt: 2200 },
  { name: "Page C", uv: 500, pv: 2600, amt: 2500 },
  { name: "Page D", uv: 450, pv: 2100, amt: 2300 },
  { name: "Page E", uv: 350, pv: 1900, amt: 2000 },
  { name: "Page F", uv: 600, pv: 2700, amt: 2700 },
];

const LineGraph = () => {
  return (
    <ClientOnly>
      <div>
        <LineChart
          width={400}
          height={300}
          data={data}
          margin={{ top: 5, right: 20, bottom: 5, left: 0 }}
        >
          <Line type="monotone" dataKey="uv" stroke="#8884d8" />
          <CartesianGrid stroke="#ccc" strokeDasharray="5 5" />
          <XAxis dataKey="name" />
          <YAxis />
          <Tooltip />
        </LineChart>
      </div>
    </ClientOnly>
  );
};

export default LineGraph;
