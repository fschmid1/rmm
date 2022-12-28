<script lang="ts">
import {
  onMount,
  onDestroy
} from "svelte";
import {
  callDeviceFunction
} from "../http";
import {
  ws
} from "../stores";
import type {
  Device
} from "../types";
import ChartStreaming from 'chartjs-plugin-streaming';
import 'chartjs-adapter-date-fns';
import {
  enUS
} from 'date-fns/locale';

import {
  Chart,
} from 'chart.js/auto';

export let device: Device;

let dataQueue = [];
let canvasElement: HTMLCanvasElement;
let chart: Chart;

let data = {
  labels: [],
  datasets: [{
      label: 'CPU',
      data: [],
      borderColor: 'rgb(54, 162, 235)',
      fill: false,
    },
    {
      label: 'Memory',
      data: [],
      borderColor: 'rgb(255, 99, 132)',
      fill: false,
    },
  ],
};
let options: any = {
  responsive: true,
  scales: {
    y: {
      beginAtZero: true,
      max: 100,
			ticks: {
				callback: value => `${value}%`,
			}
    },
    x: {
      type: 'realtime',
      adapters: {
        date: {
          locale: enUS,
        },
      },
      realtime: {
        duration: 20000,
        onRefresh: chart => {
          const data = dataQueue.pop();
          if (!data) return;
          chart.data.datasets[0].data.push({
            x: Date.now(),
            y: data.cpu,
          });
          chart.data.datasets[1].data.push({
            x: Date.now(),
            y: data.memory,
          });
        }
      }
    }
  }
};
onMount(async () => {
  Chart.register(ChartStreaming);
  chart = new Chart(canvasElement, {
    type: 'line',
    data: data,
    options: options,
  });

  $ws.on('usage', (data) => {
    dataQueue.push(data);
  });
  await callDeviceFunction(device.deviceID, 'usage-start');
});

onDestroy(async () => {
  $ws.off('usage');
	const interval_id = window.setInterval(function(){}, Number.MAX_SAFE_INTEGER);
	for (let i = 1; i < interval_id; i++) {
 		window.clearInterval(i);
	}
  await callDeviceFunction(device.deviceID, 'usage-stop');
})
</script>

<canvas bind:this={canvasElement}></canvas>
