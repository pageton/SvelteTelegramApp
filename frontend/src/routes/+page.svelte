<script lang="ts">
  import { onMount } from "svelte";
  import axios from "axios";

  interface WebAppUser {
    id: number;
    first_name: string;
    last_name?: string;
    username?: string;
    photo_url?: string;
  }

  let isHashValid: boolean | null = null;
  let userData: WebAppUser | null = null;

  const validateHash = async () => {
    try {
      const response = await axios.post("http://localhost:3000/validate", {
        hash: window.Telegram.WebApp.initData,
      });

      isHashValid = response.data.ok === true;
    } catch (error) {
      console.error("Failed to validate hash:", error);
      isHashValid = false;
    }
  };

  const initTelegram = () => {
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.ready();

      const user = window.Telegram.WebApp.initDataUnsafe?.user;
      if (user) {
        userData = user as WebAppUser;
      }
    }
  };

  onMount(() => {
    validateHash();

    if (document.readyState === "complete") {
      initTelegram();
    } else {
      window.addEventListener("load", initTelegram);
      return () => window.removeEventListener("load", initTelegram);
    }
  });
</script>

{#if isHashValid === null}
  <div>Loading...</div>
{:else if !isHashValid}
  <div>Invalid hash</div>
{:else}
  <main
    class="flex min-h-screen flex-col items-center justify-center bg-gradient-to-b from-[#2e026d] to-[#15162c] text-white"
  >
    <div>
      Information: <br />
      First name: {userData?.first_name}
      <br />
      User ID: {userData?.id}
    </div>
  </main>
{/if}
