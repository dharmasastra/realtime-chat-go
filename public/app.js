new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        joined: false // True if email and username have been filled in
    },

    created: function() {
        this.init();
    },

    methods: {
        init: function() {
            this.ws = new WebSocket('ws://' + window.location.host + '/ws');
            this.ws.addEventListener('message', (e) =>  {
                const msg = JSON.parse(e.data);
                this.chatContent += `<div class="chip">
                                        <img src="https://robohash.org/' + msg.username + '?size=32x32" alt="">
                                        ${msg.username}
                                      </div>
                                        ${msg.messages}<br/>`;

                const element = document.getElementById('chat-messages');
                element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
            });
        },
        send: function () {
            if (this.newMsg !== '') {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        messages: this.newMsg.replace(/<(?:.|\n)*?>/gm, '')// Strip out html
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        join: function () {
            if (!this.email) {
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.email = this.email.replace(/<(?:.|\n)*?>/gm, '');
            this.username = this.username.replace(/<(?:.|\n)*?>/gm, '');
            this.joined = true;
        },
    }
});