<script setup lang="ts">
// External component imports (assumed to exist)
import { Avatar } from "@/components/ui/Avatar";
import { Toggle } from "@/components/ui/Toggle";
import { useUserStore } from "@/stores/userStore";
import { useUserPreferences } from "@/composables/useUserPreferences";

const userStore = useUserStore();
const { preferences, toggleDarkMode } = useUserPreferences();
</script>

<template>
  <div class="card" :class="{ dark: preferences.darkMode }">
    <Avatar :src="userStore.user.avatarUrl" :alt="`${userStore.user.name}'s avatar`" />

    <div class="info">
      <h2 class="name">{{ userStore.user.name }}</h2>
      <p class="email">{{ userStore.user.email }}</p>
    </div>

    <div class="actions">
      <Toggle
        label="Dark Mode"
        :checked="preferences.darkMode"
        @change="toggleDarkMode"
      />
    </div>
  </div>
</template>

<style scoped>
.card {
  display: flex;
  align-items: center;
  padding: 1rem;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  background-color: white;
  transition: background-color 0.3s ease;
}

.card.dark {
  background-color: #2c2c2c;
  color: white;
}

.info {
  flex: 1;
  margin-left: 1rem;
}

.name {
  font-size: 1.2rem;
  font-weight: bold;
}

.email {
  font-size: 0.9rem;
  color: gray;
}

.actions {
  margin-left: auto;
}
</style>
