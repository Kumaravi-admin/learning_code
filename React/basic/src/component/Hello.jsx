// function Hello({ name, message }) {
//   return (
//     <h1>
//       {message} {name}
//     </h1>
//   );
// }

function Hello(props) {
  const { name, message } = props;
  return (
    <h1>
      {message} {name}
    </h1>
  );
}

export default Hello;
