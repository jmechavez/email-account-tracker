import Alpine from 'alpinejs'

window.Alpine = Alpine

window.userApp = function() {
    return {
        users: [],

        init() {
            this.fetchUsers();
        },

        fetchUsers() {
            fetch('http://localhost:8000/users')
                .then(response => response.json())
                .then(data => {
                    console.log('Users data:', data);
                    this.users = data;
                })
                .catch(error => {
                    console.error('Error fetching users:', error);
                    alert('Failed to load users');
                });
        }
    }
}

Alpine.start()
