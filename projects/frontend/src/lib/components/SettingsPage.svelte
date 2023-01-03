<script lang="ts">
import {
    Button,
    Li,
    List,
} from "flowbite-svelte";
import {
    Router,
    Route,
    Link
} from "svelte-navigator";
import TokenOverview from "./TokenOverview.svelte";
import Fa from 'svelte-fa/src/fa.svelte'
import {
    faArrowLeft
} from '@fortawesome/free-solid-svg-icons'
import Profile from "./Profile.svelte";
</script>
<div class="w-11/12 mx-auto mt-12">
    <Link to="/devices">
    <Button class="bg-gray-700 hover:bg-gray-600">
        <Fa icon={faArrowLeft} class="mr-2" />
        Back
    </Button>
    </Link>
    <div class="flex justify-between gap-4 mt-4">
        <div class="w-3/12 rounded-lg shadow-md border border-gray-700 ">
            <List tag="ul" class=" divide-gray-200 dark:divide-gray-700 py-2" list="none">
                <Li class="py-2 pl-4">
                    <Link to="/settings/profile">
                    	<span class="space-x-4">Profile</span>
                    </Link>
                </Li>
                <Li class="py-2 pl-4">
                    <Link to="/settings/tokens">
                    	<span class="space-x-4">Tokens</span>
                    </Link>
                </Li>
            </List>
        </div>
        <div class="flex w-9/12 flex-col">
            <div class="rounded-lg shadow-md border border-gray-700 p-4">
                <Router>
                    <Route primary={false} path="/tokens" >
                        <TokenOverview />
                    </Route>
                    <Route path="/profile" primary={false}>
						<Profile />
                    </Route>
                </Router>
            </div>
        </div>
    </div>
</div>
