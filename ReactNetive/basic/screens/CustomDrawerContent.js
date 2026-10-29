import React from "react";
import { View, Text, Button } from "react-native";

export default function CustomDrawerContent({ navigation }) {
  return (
    <View style={{ flex: 1, padding: 20, backgroundColor: "red" }}>
      <View>
        <Text>Anjesh Kumar Ravi</Text>
        <Text>anjeshravi8@yahoo.com</Text>
      </View>
      <Button title="Go to Home" onPress={() => navigation.navigate("Home")} />
      <Button
        title="Go to Settings"
        onPress={() => navigation.navigate("Settings")}
      />
    </View>
  );
}
