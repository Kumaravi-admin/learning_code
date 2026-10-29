import React from "react";
import { StyleSheet, View } from "react-native";
import Animated, {
  useSharedValue,
  useAnimatedStyle,
  withSpring,
} from "react-native-reanimated"; // Import Reanimated for animations
import { PanGestureHandler } from "react-native-gesture-handler"; // Import Gesture Handler for touch gestures

const DraggableBox = () => {
  // Shared values to store the X and Y translation values
  const translateX = useSharedValue(0);
  const translateY = useSharedValue(0);

  // Animated style to update the transform property based on shared values
  const animatedStyle = useAnimatedStyle(() => ({
    transform: [
      { translateX: translateX.value }, // Apply horizontal movement
      { translateY: translateY.value }, // Apply vertical movement
    ],
  }));

  // Gesture handler to update translation values on drag
  const onGestureEvent = ({ nativeEvent }) => {
    translateX.value = nativeEvent.translationX; // Update X translation
    translateY.value = nativeEvent.translationY; // Update Y translation
  };

  return (
    <View style={styles.container}>
      {/* Gesture handler wraps the animated view to enable dragging */}
      <PanGestureHandler onGestureEvent={onGestureEvent}>
        {/* Animated view to display the draggable box */}
        <Animated.View style={[styles.box, animatedStyle]} />
      </PanGestureHandler>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1, // Take up the full screen
    justifyContent: "center", // Center the content vertically
    alignItems: "center", // Center the content horizontally
  },
  box: {
    width: 100, // Width of the box
    height: 100, // Height of the box
    backgroundColor: "#61dafb", // Light blue background
    borderRadius: 10, // Rounded corners
  },
});

export default DraggableBox; // Export the component
