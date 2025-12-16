<template>
    <div class="group-create-container">
      <h1 class="page-title">Create Group</h1>
      <div class="form-group">
        <label for="group-name">Group Name:</label>
        <input
          id="group-name"
          v-model="groupName"
          class="input-field"
          type="text"
          placeholder="Enter group name"
        />
      </div>
      <form @submit.prevent="searchUsers" class="search-form">
        <input
          id="username"
          v-model="query"
          class="search-box"
          type="text"
          placeholder=" Search by username"
        />


        <button class="search-button" type="submit">Search</button>
      </form>
      <div v-if="error" class="error-box">
        {{ error }}
      </div>
      <div v-if="loading">
        <LoadingSpinner />
      </div>
      <div v-if="!loading && showResults" class="results-section">
        <h2 class="results-title">Search Results:</h2>
        <template v-if="users.length > 0">
          <div v-for="user in users" :key="user.id" class="user-card">
            <h5 class="user-name">@{{ user.name }}</h5>
            <button
              class="add-button"
              @click="addUserToGroup(user)"
              :disabled="isUserAdded(user)"
            >
              {{ isUserAdded(user) ? 'Added' : 'Add' }}
            </button>
          </div>
        </template>
        <template v-else>
          <p class="no-results">No users found matching "{{ lastQuery }}"</p>
        </template>
      </div>
      <div v-if="selectedUsers.length > 0" class="selected-users-section">
        <h2 class="selected-title">Selected Members:</h2>
        <div class="selected-users">
          <span v-for="user in selectedUsers" :key="user.id" class="selected-user">
            @{{ user.name }}
            <button @click="removeUserFromGroup(user)" class="remove-button">Remove</button>
          </span>
        </div>
      </div>
      <div class="form-group">
        <label for="group-image">Group Image:</label>
        <input
          id="group-image"
          ref="fileInput"
          type="file"
          @change="handleImageUpload"
          accept="image/*"
        />
        <img v-if="previewImage" :src="previewImage" class="preview-image" alt="Group Image Preview"/>
      </div>
      <button type="button" class="create-button" @click="createGroup">
        Create Group
      </button>

    </div>
  </template>

  <script>
  import axios from "../services/axios";
  import LoadingSpinner from "../components/LoadingSpinner.vue";
  
  export default {
    name: "GroupCreateView",
    components: {
      LoadingSpinner,
    },
    data() {
      const token = localStorage.getItem("token");
      if (!token) {
        this.$router.push({ path: "/" });
        return;
      }
      return {
        groupName: "",
        query: "",
        lastQuery: "",
        users: [],
        loading: false,
        showResults: false,
        error: "",
        selectedUsers: [],
        previewImage: null,
        file: null,
      };
    },
    computed: {
      canCreateGroup() {
        return (
          this.groupName.trim() !== "" &&
          this.selectedUsers.length > 0
        );
      },
    },
    methods: {
      async searchUsers() {
        if (!this.query.trim()) {
          this.error = "Please enter a valid search query.";
          this.showResults = false;
          return;
        }
        this.loading = true;
        this.error = "";
        this.users = [];
        this.showResults = false;
        try {
          const response = await axios.get(`/search`, {
            params: { username: this.query },
          });
          this.users = response.data.filter(user => user.id !== localStorage.getItem("token"));
          this.lastQuery = this.query;
          this.showResults = true;
        } catch (err) {
          const status = err.response?.status;
          const reason = err.response?.data?.message || "Failed to fetch users.";
          this.error = `Status ${status}: ${reason}`;
        } finally {
          this.loading = false;
        }
      },
      addUserToGroup(user) {
        if (!this.isUserAdded(user)) {
          this.selectedUsers.push(user);
        }
      },
      isUserAdded(user) {
        return this.selectedUsers.some((u) => u.id === user.id);
      },
      removeUserFromGroup(user) {
        this.selectedUsers = this.selectedUsers.filter((u) => u.id !== user.id);
      },
      handleImageUpload(event) {
        const file = event.target.files[0];
        if (file) {
          this.file = file;
          const reader = new FileReader();
          reader.onload = (e) => {
            this.previewImage = e.target.result;
          };
          reader.readAsDataURL(file);
        } else {
          this.previewImage = null;
          this.file = null;
        }
      },
      async createGroup() {
        console.log("createGroup clicked");
        if (this.groupName.trim() === "") {
          alert("Please enter a group name");
          return;
        }
        if (this.selectedUsers.length === 0) {
          alert("Please add at least one member");
          return;
        }
        this.loading = true;
        this.error = "";
        try {
          const formData = new FormData();
          formData.append("name", this.groupName);
          formData.append("members", JSON.stringify([...this.selectedUsers.map(u => u.id), localStorage.getItem("token")
          ])
        );
        if (this.file) {
          formData.append("image", this.file);
        } else {
          const res = await fetch("/default-group.png");
          if (!res.ok) throw new Error("Default image not found");

          const blob = await res.blob();
          formData.append("image", blob, "default.png");
        }
        await axios.post(`/groups`, formData); 
        
        alert("Group created successfully!");
        this.$router.push(`/home`);
        
        } catch (err) {
          console.error(err);
          this.error = "Failed to create group";
        } finally {
          this.loading = false;
        }
      },
    },
  };
  </script>

  <style scoped>

.group-create-container h1,
.group-create-container h2,
.group-create-container h3,
.group-create-container h4,
.group-create-container h5,
.group-create-container p,
.group-create-container span {
  color: #f2b7c7ff;
  text-shadow: 0 0 2px #6b4f4f;
}
  .group-create-container {
    --form-width: 420px;
    text-align: center;
    padding: 20px;
    max-width: 600px;
    margin: 120px auto;
    color: #f2b7c7ff;
    text-shadow: 0 0 2px #6b4f4f;
  }

  .page-title {
  color: #f2b7c7ff;
  text-shadow: 0 0 2px #6b4f4f;
  }

  .results-title {
    font-size: 16px;
    margin-bottom: 8px;
    color: #f2b7c7ff;
    text-shadow: 0 0 2px #6b4f4f;
  }
  label {
    color: #e0b4c0ff;
  }

  .user-card {
    display: flex;
    justify-content: space-between;
    align-items: center;

    padding: 10px 12px;
    margin-top: 8px;

    background-color: transparent;
    border-radius: 8px;
    border: none;
  }
  .user-card:hover {
    background-color: #f0c6d262;
  }

  .form-group {
    margin-bottom: 24px;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .form-group label {
    margin-bottom: 8px;
  }



  .search-button {
    width: 100px;          
    flex-shrink: 0;
  }

  .input-field {
    width: var(--form-width);
    max-width: 100%;
    padding: 12px;
    background-color: #3a3a3c;  
    color: #ffdde6;
    border: 1px solid #6b4f4f;
    border-radius: 8px;          
  }

  .search-box,
  input[type="file"] {
    width: var(--form-width);
    max-width: 100%;
    background-color: #3a3a3c;  /* DARK like before */
    color: #ffdde6;
    border-radius: 8px;
    border: 1px solid #6b4f4f;
  }

  .input-field::placeholder,
  .search-box::placeholder {
    color: #d9a7b5;
  }

  .input-field:focus,
  .search-box:focus {
    outline: none;
    border-color: #ff779b;   /* subtle pink */
    box-shadow: none;        /* IMPORTANT */
  }

  .search-form {
    width: var(--form-width);
    max-width: 100%;
    margin: 0 auto 24px;
    display: flex;
    gap: 12px;
  }
  
  .search-button,
  .add-button,
  .remove-button {
    background-color: #ff779b;
    color: #ffffff;
    border: none;
    border-radius: 8px;
    padding: 10px 16px;
    font-weight: 600;
    cursor: pointer;
  }
    
  .create-button {
    background-color: #ff779b;
    color: #ffffff;
    border: none;
    border-radius: 10px;
    padding: 12px 24px;
    font-weight: 600;
    font-size: 16px;
    cursor: pointer;
    margin-top: 20px;
  } 

  .create-button:hover {
    background-color: #ff8fb1;
  }

  .create-button:disabled {
    background-color: #555;
    cursor: not-allowed;
  }

  .user-name {
    color: #3b2b2b; /* тёмно-серый для @ajaja */
  }

  .selected-user {
    color: #3b2b2b;
  }



  .search-button:hover,
  .add-button:hover,
  .remove-button:hover {
    background-color: #ff8fb1;
  }
    
  button:disabled {
    background-color: #555;
    cursor: not-allowed;
  }

  .error-box {
    background-color: #3a1f28;
    color: #ffb3c6;
    border: 1px solid #6b4f4f;
    border-radius: 8px;
    padding: 12px;
    margin: 20px 0;
  }



    .selected-user {
      background-color: #696465ff;
      color: #3b2a2a;
      border-radius: 8px;
      padding: 6px 12px;
      margin: 5px;
    }

    .preview-image {
      max-width: 200px;
      max-height: 200px;
      margin-top: 12px;
      border-radius: 10px;
      border: 1px solid #6b4f4f;
    }

    .results-section {
      width: var(--form-width);
      max-width: 100%;
      margin: 14px auto;
      background-color: #2f2f33;   /* dark card */
      border: 1px solid #6b4f4f;
      border-radius: 12px;
      padding: 10px;
    }
    .results-wrapper {
      min-height: 160px;   /* smaller than before */
    }

    input[type="file"]::file-selector-button {
      background-color: #ebb8cdff;   /* light grey */
      color: #2b2b2b;
      border-color: #ff779b; 
      border-radius: 6px;
      padding: 4px 11px;
      margin-right: 11px;
      cursor: pointer;
    }
    </style>