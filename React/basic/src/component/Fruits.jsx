import Fruit from "./Fruit";

export default function Fruits() {
  const fruits = [
    { name: "Apple", color: "Red", weight: 180, price: 120 },
    { name: "Banana", color: "Yellow", weight: 120, price: 40 },
    { name: "Orange", color: "Orange", weight: 150, price: 60 },
    { name: "Grapes", color: "Purple", weight: 50, price: 90 },
    { name: "Mango", color: "Green", weight: 200, price: 150 },
  ];

  return (
    <>
      {fruits.map((fruit) => (
        <Fruit
          name={fruit.name}
          color={fruit.color}
          weight={fruit.weight}
          price={fruit.price}
        />
      ))}
    </>
  );
}
