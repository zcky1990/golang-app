import countDown from "../../static/javascript/components/CountDown.vue";

export default {
  title: "Components/CountDown",
  component: countDown,
  tags: ["autodocs"],
  parameters: {
    docs: {
      description: {
        component: 'A countdown timer component that counts down to a specified date. Customize the title and the end date using props.'
      }
    }
  },
  render: (args) => ({
    components: {
      countDown,
    },
    setup() {
      return {
        ...args,
      };
    },
    template:
      '<countDown :position="position" :targetDate="targetDate" />',
  }),
  argTypes: {
    position: { 
      options: ['center', 'left'],
      control: { type: 'select' }, },
    targetDate: { control: "text" },
  },
};

export const Center = {
  parameters: {
    docs: {
      description: {
        story: 'Example how countdown component renderer when position is center'
      },
      source: {
        code: `<countDown position="center" targetDate="2025-03-25" />`
      },
    }
  },
  args: {
    position: "center",
    targetDate: "2025-03-25"
  }
}

export const Left = {
  parameters: {
    docs: {
      description: {
        story: 'Example how countdown component renderer when position is left'
      },
      source: {
        code: `<countDown position="left" targetDate="2025-03-25" />`
      },
    }
  },
  args: {
    position: "left",
    targetDate: "2025-03-25"
  },
};


export const Right = {
  parameters: {
    docs: {
      description: {
        story: 'Example how countdown component renderer when position is right'
      },
      source: {
        code: `<countDown position="right" targetDate="2025-03-25" />`
      },
    }
  },
  args: {
    position: "right",
    targetDate: "2025-03-25"
  },
};
