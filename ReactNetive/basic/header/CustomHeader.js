import { View, Text, TouchableOpacity, StyleSheet, Button } from "react-native";

function CustomHeader({ title, onPressBack }) {
  return (
    <View style={styles.header}>
      <Button title="Back" onPress={onPressBack} />
      <Text style={styles.title}>{title}</Text>
    </View>
  );
}
const styles = StyleSheet.create({
  header: {
    height: 160,
    alignItems: "center",
    justifyContent: "center",
    padding: 10,
    backgroundColor: "black",
    elevation: 10,
  },
  title: {
    fontSize: 20,
    marginLeft: 10,
    color: "white",
  },
});

export default CustomHeader;
