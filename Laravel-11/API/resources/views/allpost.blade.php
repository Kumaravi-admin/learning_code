<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
        <h1>All Post Page</h1>
        <button id="logoutButton">Logout</button>

        <div id="container"></div>


        <div id="modelBox" style="z-index: 1;display:none">
          
            <button id="buttonClose"> Close</button>
    
        </div>
        
        
      

<script>
    document.querySelector('#logoutButton').addEventListener('click',function(){
        const token = localStorage.getItem('api_token');

        fetch('http://localhost:8000/api/logout',{
            method:'POST',
            headers:{
                'Authorization':`Bearer ${token}`,
            }
        }).then(responce => responce.json())
          .then(data => {
            console.log(data);
            localStorage.removeItem('api_token');
            window.location.href = 'http://localhost:8000/'; 
        });
    });


    
    function loadData(){
            const token = localStorage.getItem('api_token');

            fetch('http://localhost:8000/api/posts',{
                method:'GET',
                headers:{
                    'Authorization':`Bearer ${token}`,
                }
            }).then(responce => responce.json())
                .then(data => {
                    console.log(data);

                    const container = document.querySelector('#container');
                    const allPosts = data.data.posts;
                    var postTable = `<table style="border: 0.5px solid black;">
                                        <tr>
                                            <th>Image</th>
                                            <th>Title</th>
                                            <th>Description</th>
                                            <th>View</th>
                                            <th>Update</th>
                                            <th>Delete</th>
                                        </tr>`;
                    allPosts.forEach(post => {
                        postTable += `<tr>
                                    <td><img src='/uploads/${post.image}' width='40px'></td>
                                    <td><h6>${post.title}</h6></td>
                                    <td><p>${post.description}</p></td>
                                    <td><button id="view" data-postid="${post.id}" class="viewButton">View</button></td>
                                    <td><button id="Update">Update</button></td>
                                    <td><button id="delete">Delete</button></td>
                                  </tr>`;
                        
                    });


                    postTable += `</table>`;
                    container.innerHTML = postTable;
                
                alert('code working');

                  document.querySelectorAll('.viewButton').forEach(button =>{
                    button.addEventListener('click', event =>{
                      const postId = event.target.getAttribute('data-postid');
                      viewPost(postId);

                    });
                });
        

         });
        }

        loadData();
    

        function viewPost(postId){

            const token = localStorage.getItem('api_token');

            fetch(`http://localhost:8000/api/posts/${postId}`,{
                method:'GET',
                headers:{
                    'Authorization':`Bearer ${token}`,
                    'Content-Type' : 'application/json'
                }

            }).then(responce => responce.json())
            .then(data =>{
                var post = data.data.post;
                var modelBox = document.querySelector('#modelBox');
                modelBox.innerHTML = "";

                modelBox.innerHTML = `
                Tile : ${post.title}
                <br>
                Description : ${post.description}
                <br>
                <img width="100px" src="http://localhost:8000/uploads/${post.image}">       
                `;
                modelBox.style.display = 'block';
            console.log(data); 
            });
    

        }
        
        
        
        
        
        



</script>
</body>
</html>