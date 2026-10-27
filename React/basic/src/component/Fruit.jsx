export default function Fruit({ name, color, weight, price }) {
  return (
    <>
      {price > 50 ? (
        <p>
          {name} - {color} - {weight} - rs {price}
        </p>
      ) : (
        ""
      )}
    </>
  );
}
