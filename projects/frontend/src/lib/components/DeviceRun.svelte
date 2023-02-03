<script lang="ts">
    import { Input } from 'flowbite-svelte';
    import type { Device } from '../../types';
    import { callDeviceFunction } from '../helper/http';

    export let device: Device;
    let commands: string[] = [];
    let commandIndex = -1;
    let command = '';
    let output = '';
    let outputElement: HTMLDivElement;
    let controllPressed = false;

    const trimOutput = () => {
        while (output.endsWith('<br /><br />')) {
            output = output.slice(0, -6);
        }
    };

    const scrollToBottom = async (node, data = '') => {
        if (commands.length > 0) {
            output += new Array(18 - data.split('\n').length).join('<br />');
        }
        const options: ScrollToOptions = { top: node.scrollHeight, behavior: 'auto' };
        setTimeout(() => {
            node.scroll(options);
        }, 0);
    };

    const handleKey = async (event: any) => {
        if (event.key === 'Enter') {
            trimOutput();
            callDeviceFunction<string>(device.deviceID, 'run', command).then((response) => {
                output += response.data.replace(/\n/g, '<br />');
                scrollToBottom(outputElement, response.data);
            });
            commands = [command, ...commands];
            command = '';
            commandIndex = -1;
        } else if (event.key === 'ArrowUp') {
            event.preventDefault();
            if (commandIndex < commands.length - 1 || commands.length === -1) {
                commandIndex = commandIndex + 1;
                command = commands[commandIndex];
            }
        } else if (event.key === 'ArrowDown') {
            event.preventDefault();
            if (commandIndex > 0) {
                commandIndex = commandIndex - 1;
                command = commands[commandIndex];
            } else if (commandIndex === 0) {
                commandIndex = -1;
                command = '';
            }
        } else if (event.key === 'Control') {
            controllPressed = true;
        } else if (event.key === 'c' && controllPressed) {
            event.preventDefault();
            command = '';
        } else if (event.key === 'l' && controllPressed) {
            event.preventDefault();
            trimOutput();
            scrollToBottom(outputElement);
        }
    };
</script>

<div class="mb-2 h-96 overflow-y-scroll" bind:this={outputElement}>
    {@html output}
</div>

<Input
    type="text"
    bind:value={command}
    id="command"
    placeholder=""
    required
    on:keydown={handleKey}
    on:keyup={(event) => {
        if (event.key === 'Control') {
            controllPressed = false;
        }
    }}
/>
